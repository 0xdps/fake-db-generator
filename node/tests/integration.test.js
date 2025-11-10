/**
 * Simple integration tests for fakestack
 * Run with: node tests/integration.test.js
 */

const assert = require('assert');
const fs = require('fs');
const path = require('path');
const { fakestack, getBinaryPath } = require('../dist/index');

async function runTests() {
  console.log('🧪 Running fakestack integration tests...\n');

  // Test 1: Import
  console.log('1. Testing imports...');
  assert.ok(fakestack, 'fakestack function should be exported');
  assert.ok(getBinaryPath, 'getBinaryPath function should be exported');
  console.log('   ✓ Imports successful\n');

  // Test 2: Binary detection
  console.log('2. Testing binary detection...');
  const binaryPath = getBinaryPath();
  assert.ok(fs.existsSync(binaryPath), `Binary should exist at ${binaryPath}`);
  console.log(`   ✓ Binary found at: ${binaryPath}\n`);

  // Test 3: Download schema
  console.log('3. Testing schema download...');
  const schemaPath = path.join(process.cwd(), 'schema.json');
  if (fs.existsSync(schemaPath)) {
    fs.unlinkSync(schemaPath);
  }
  const exitCode1 = await fakestack(['-d', '.']);
  assert.strictEqual(exitCode1, 0, 'Schema download should succeed');
  assert.ok(fs.existsSync(schemaPath), 'schema.json should be created');
  console.log('   ✓ Schema downloaded successfully\n');

  // Test 4: Create and populate database
  console.log('4. Testing database creation...');
  const exitCode2 = await fakestack(['-c', '-p', '-f', 'schema.json']);
  assert.strictEqual(exitCode2, 0, 'Database creation should succeed');
  assert.ok(fs.existsSync('test.db'), 'test.db should be created');
  const stats = fs.statSync('test.db');
  assert.ok(stats.size > 0, 'Database should not be empty');
  console.log(`   ✓ Database created (${(stats.size / 1024).toFixed(1)} KB)\n`);

  // Test 5: Error handling
  console.log('5. Testing error handling...');
  try {
    await fakestack([]);
    assert.fail('Should have thrown error for invalid args');
  } catch (err) {
    assert.ok(err.message.includes('exited with code'), 'Should throw error with exit code');
    console.log('   ✓ Errors handled correctly\n');
  }

  console.log('✅ All tests passed!');
  
  // Cleanup
  if (fs.existsSync('schema.json')) fs.unlinkSync('schema.json');
  if (fs.existsSync('test.db')) fs.unlinkSync('test.db');
}

// Run tests
if (require.main === module) {
  runTests().catch(err => {
    console.error('\n❌ Tests failed:', err.message);
    process.exit(1);
  });
}

module.exports = { runTests };
