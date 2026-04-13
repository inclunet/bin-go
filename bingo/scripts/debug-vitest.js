import { startVitest } from 'vitest/node';

(async () => {
  console.log('Starting vitest programmatically...');
  const ctx = await startVitest('test', [], { reporters: ['basic'] });
  console.log('Finished with status', ctx); 
})();