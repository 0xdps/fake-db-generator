/**
 * Fakestack - High-Performance Database Generator (TypeScript Wrapper)
 * 
 * This package wraps the Go core binary for blazing-fast database generation.
 * Uses external process execution for simplicity and cross-platform compatibility.
 */

import { spawn, SpawnOptions } from 'child_process';
import { platform, arch } from 'os';
import { join } from 'path';
import { chmodSync, existsSync } from 'fs';

/**
 * Get the path to the platform-specific binary
 */
export function getBinaryPath(): string {
  // Platform mapping
  const platformMap: Record<string, string> = {
    'linux': 'linux',
    'darwin': 'darwin',
    'win32': 'windows'
  };
  
  // Architecture mapping
  const archMap: Record<string, string> = {
    'x64': 'amd64',
    'arm64': 'arm64'
  };
  
  const osName = platformMap[platform()] || platform();
  const archName = archMap[arch()] || 'amd64';
  
  // Construct binary name
  let binaryName = `fakestack-${osName}-${archName}`;
  if (osName === 'windows') {
    binaryName += '.exe';
  }
  
  // Find binary path - CommonJS module
  const binDir = join(__dirname, '..', 'bin');
  const binaryPath = join(binDir, binaryName);
  
  if (!existsSync(binaryPath)) {
    throw new Error(
      `Binary not found: ${binaryPath}\n` +
      `Platform: ${osName}-${archName}\n` +
      `Please report this issue at: https://github.com/0xdps/fake-stack/issues`
    );
  }
  
  return binaryPath;
}

/**
 * Execute the fakestack Go binary with given arguments
 */
export function runFakestack(args: string[]): Promise<number> {
  return new Promise((resolve, reject) => {
    try {
      const binary = getBinaryPath();
      
      // Make executable on Unix systems
      if (platform() !== 'win32') {
        try {
          chmodSync(binary, 0o755);
        } catch (err) {
          // Already executable or no permission - ignore
        }
      }
      
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
