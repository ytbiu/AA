'use client';

import { useState } from 'react';
import { AdminPanel } from '@/components/AdminPanel';

export default function AdminPage() {
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleAdminAction = async (action: string, params: any) => {
    setLoading(true);
    setResult(null);

    try {
      // This would connect to actual admin APIs
      // For now we'll simulate the response
      let response;

      switch(action) {
        case 'addRelayer':
          response = { success: true, action: 'addRelayer', params };
          break;
        case 'removeRelayer':
          response = { success: true, action: 'removeRelayer', params };
          break;
        case 'setFeeRate':
          response = { success: true, action: 'setFeeRate', params };
          break;
        case 'setOracle':
          response = { success: true, action: 'setOracle', params };
          break;
        default:
          response = { success: false, error: 'Unknown action' };
      }

      // In a real implementation, we would call the API like:
      // const apiResponse = await fetch('/api/admin', {
      //   method: 'POST',
      //   headers: { 'Content-Type': 'application/json' },
      //   body: JSON.stringify({ action, ...params })
      // });
      // response = await apiResponse.json();

      setResult(response);
    } catch (error) {
      console.error('Error performing admin action:', error);
      setResult({ error: `Failed to perform ${action} action` });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Admin Panel</h1>

      <div className="mb-6 p-4 bg-red-100 border border-red-400 rounded-lg">
        <h2 className="text-lg font-semibold mb-2">Authentication Warning</h2>
        <p>
          This admin panel currently has no authentication implemented.
          In a production environment, you would need to implement proper
          authentication and authorization before allowing access to these
          administrative functions.
        </p>
      </div>

      <div className="mb-6">
        <AdminPanel onAction={handleAdminAction} loading={loading} />
      </div>

      {result && (
        <div className="mt-6 p-4 bg-gray-100 rounded-lg">
          <h3 className="font-semibold mb-2">Result:</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}