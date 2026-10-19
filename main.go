"use client";
import { useState, useEffect } from "react";

export default function Home() {
  const [result, setResult] = useState<any>(null);
  const [status, setStatus] = useState("");

  // Generic fetch function to talk to Go
  const callAPI = async (endpoint: string, body: any = null) => {
    setResult("Loading...");
    try {
      const res = await fetch(`http://localhost:8080${endpoint}`, {
        method: body ? "POST" : "GET",
        headers: { "Content-Type": "application/json" },
        body: body ? JSON.stringify(body) : null,
      });
      const data = await res.json();
      setResult(data);
    } catch (err) {
      setResult({ error: "Failed to connect to Go" });
    }
  };

  // Scenario 4 runs automatically on load
  useEffect(() => {
    callAPI("/api/status");
  }, []);

  return (
    <main className="min-h-screen bg-gray-900 text-white p-8 font-mono">
      <h1 className="text-3xl font-bold mb-6 text-blue-400">Go + Next.js Scenarios</h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8">
        {/* Scenario 1 */}
        <div className="bg-gray-800 p-4 rounded">
          <h2 className="font-bold mb-2">1. Text (Hello)</h2>
          <button onClick={() => callAPI("/api/hello", { name: "SableDkaze" })} className="bg-blue-600 px-4 py-2 rounded">
            Say Hello
          </button>
        </div>

        {/* Scenario 2 */}
        <div className="bg-gray-800 p-4 rounded">
          <h2 className="font-bold mb-2">2. Math (Addition)</h2>
          <button onClick={() => callAPI("/api/add", { A: 15, B: 27 })} className="bg-blue-600 px-4 py-2 rounded">
            Add 15 + 27
          </button>
        </div>

        {/* Scenario 3 */}
        <div className="bg-gray-800 p-4 rounded">
          <h2 className="font-bold mb-2">3. Arrays (Stats)</h2>
          <button onClick={() => callAPI("/api/stats", { numbers: [10, 20, 30, 40] })} className="bg-blue-600 px-4 py-2 rounded">
            Calculate Stats
          </button>
        </div>

        {/* Scenario 4 */}
        <div className="bg-gray-800 p-4 rounded">
          <h2 className="font-bold mb-2">4. Status (GET)</h2>
          <button onClick={() => callAPI("/api/status")} className="bg-blue-600 px-4 py-2 rounded">
            Check Server
          </button>
        </div>

        {/* Scenario 5 */}
        <div className="bg-gray-800 p-4 rounded">
          <h2 className="font-bold mb-2">5. Errors (Login)</h2>
          <button onClick={() => callAPI("/api/login", { password: "wrong" })} className="bg-red-600 px-4 py-2 rounded">
            Login (Wrong)
          </button>
          <button onClick={() => callAPI("/api/login", { password: "secret123" })} className="bg-green-600 px-4 py-2 rounded ml-2">
            Login (Right)
          </button>
        </div>
      </div>

      {/* The Result Box */}
      <div className="bg-black p-6 rounded border border-gray-700">
        <h2 className="text-gray-400 mb-2">Response from Go:</h2>
        <pre className="text-green-400 text-sm overflow-auto">
          {JSON.stringify(result, null, 2)}
        </pre>
      </div>
    </main>
  );
}