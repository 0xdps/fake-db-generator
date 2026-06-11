/**
 * Fakestack - High-Performance Database Generator (TypeScript Wrapper)
 * 
 * This package wraps the Go core binary for blazing-fast database generation.
 * Downloads binary on first run and auto-updates when new versions are available.
 */

import { spawn, SpawnOptions } from 'child_process';
import { platform, arch, homedir } from 'os';
import { join } from 'path';
import { chmodSync, existsSync, mkdirSync, writeFileSync, readFileSync, createWriteStream } from 'fs';
import { get as httpsGet } from 'https';

const GITHUB_REPO = '0xdps/fake-stack';
const CACHE_DIR = join(homedir(), '.fakestack', 'bin');
const VERSION_FILE = join(homedir(), '.fakestack', 'version.txt');

/**
 * Get platform-specific binary information
 */
function getPlatformInfo(): { os: string; arch: string; binaryName: string } {
  const platformMap: Record<string, string> = {
    'linux': 'linux',
    'darwin': 'darwin',
    'win32': 'windows'
  };
  
  const archMap: Record<string, string> = {
    'x64': 'amd64',
    'arm64': 'arm64'
  };
  
  const osName = platformMap[platform()] || platform();
  const archName = archMap[arch()] || 'amd64';
  
  let binaryName = `fakestack-${osName}-${archName}`;
  if (osName === 'windows') {
    binaryName += '.exe';
  }
  
  return { os: osName, arch: archName, binaryName };
}

/**
 * Fetch latest version from GitHub releases
 */
async function getLatestVersion(): Promise<string> {
  return new Promise((resolve, reject) => {
    const options = {
      hostname: 'api.github.com',
      path: `/repos/${GITHUB_REPO}/releases/latest`,
      headers: { 'User-Agent': 'fakestack-npm' }
    };
    
    httpsGet(options, (res) => {
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => {
        try {
          const json = JSON.parse(data);
          const version = json.tag_name?.replace(/^v/, '') || '1.2.0';
          resolve(version);
        } catch (err) {
          reject(new Error('Failed to parse GitHub API response'));
        }
      });
    }).on('error', reject);
  });
}

/**
 * Download binary from GitHub releases
 */
async function downloadBinary(version: string, binaryName: string, targetPath: string): Promise<void> {
  return new Promise((resolve, reject) => {
    console.log(`📦 Downloading fakestack v${version}...`);
    
    const url = `https://github.com/${GITHUB_REPO}/releases/download/v${version}/${binaryName}`;
    
    httpsGet(url, (res) => {
      if (res.statusCode === 302 || res.statusCode === 301) {
        // Follow redirect
        httpsGet(res.headers.location!, (redirectRes) => {
          const file = createWriteStream(targetPath);
          redirectRes.pipe(file);
          file.on('finish', () => {
            file.close();
            chmodSync(targetPath, 0o755);
            console.log('✓ Download complete!');
            resolve();
          });
        }).on('error', reject);
      } else {
        const file = createWriteStream(targetPath);
        res.pipe(file);
        file.on('finish', () => {
          file.close();
          chmodSync(targetPath, 0o755);
          console.log('✓ Download complete!');
          resolve();
        });
      }
    }).on('error', reject);
  });
}

/**
 * Get or download the binary, checking for updates
 */
async function ensureBinary(): Promise<string> {
  const { binaryName } = getPlatformInfo();
  const binaryPath = join(CACHE_DIR, binaryName);
  
  // Create cache directory if it doesn't exist
  if (!existsSync(CACHE_DIR)) {
    mkdirSync(CACHE_DIR, { recursive: true });
  }
  
  // Check if binary exists
  const binaryExists = existsSync(binaryPath);
  
  // Get latest version
  let latestVersion: string;
  try {
    latestVersion = await getLatestVersion();
  } catch (err) {
    // If we can't check version but binary exists, use it
    if (binaryExists) {
      return binaryPath;
    }
    throw new Error('Failed to check for latest version and no local binary found');
  }
  
  // Check local version
  let localVersion = '';
  if (existsSync(VERSION_FILE)) {
    try {
      localVersion = readFileSync(VERSION_FILE, 'utf8').trim();
    } catch (err) {
      // Ignore
    }
  }
  
  // Download if missing or outdated
  if (!binaryExists || localVersion !== latestVersion) {
    if (binaryExists && localVersion !== latestVersion) {
      console.log(`🔄 Updating from v${localVersion} to v${latestVersion}...`);
    }
    
    await downloadBinary(latestVersion, binaryName, binaryPath);
    
    // Save version
    writeFileSync(VERSION_FILE, latestVersion);
  }
  
  return binaryPath;
}

/**
 * Get the path to the platform-specific binary (legacy, for compatibility)
 */
export function getBinaryPath(): string {
  const { binaryName } = getPlatformInfo();
  return join(CACHE_DIR, binaryName);
}

/**
 * Execute the fakestack Go binary with given arguments
 */
export async function runFakestack(args: string[]): Promise<number> {
  return new Promise(async (resolve, reject) => {
    try {
      // Ensure binary is downloaded and up-to-date
      const binary = await ensureBinary();
      
      // Spawn options
      const options: SpawnOptions = {
        stdio: 'inherit', // Stream output directly to parent's stdio
        cwd: process.cwd()
      };
      
      // Execute the binary
      const child = spawn(binary, args, options);
      
      child.on('close', (code) => {
        resolve(code || 0);
      });
      
      child.on('error', (err) => {
        reject(new Error(`Failed to execute binary: ${err.message}`));
      });
      
    } catch (err) {
      reject(err);
    }
  });
}

/**
 * CLI entry point
 */
export async function cli(): Promise<void> {
  try {
    const args = process.argv.slice(2);
    const exitCode = await runFakestack(args);
    process.exit(exitCode);
  } catch (err) {
    if (err instanceof Error) {
      console.error(`Error: ${err.message}`);
    } else {
      console.error('Unexpected error:', err);
    }
    process.exit(1);
  }
}

/**
 * Programmatic API
 */
export interface FakestackOptions {
  createTables?: boolean;
  populateData?: boolean;
  schemaFile?: string;
  downloadSchema?: string;
}

/**
 * Run fakestack with typed options or raw args array
 */
export async function fakestack(optionsOrArgs: FakestackOptions | string[]): Promise<number> {
  let args: string[] = [];
  
  if (Array.isArray(optionsOrArgs)) {
    // Raw args array
    args = optionsOrArgs;
  } else {
    // Typed options
    const options = optionsOrArgs;
    
    if (options.createTables) {
      args.push('-c');
    }
    
    if (options.populateData) {
      args.push('-p');
    }
    
    if (options.schemaFile) {
      args.push('-f', options.schemaFile);
    }
    
    if (options.downloadSchema) {
      args.push('-d', options.downloadSchema);
    }
  }
  
  const exitCode = await runFakestack(args);
  
  if (exitCode !== 0) {
    throw new Error(`Fakestack exited with code ${exitCode}`);
  }
  
  return exitCode;
}

// Default export
export default {
  fakestack,
  runFakestack,
  getBinaryPath,
  cli
};
