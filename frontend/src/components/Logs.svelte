<script lang="ts">
  import { onMount } from "svelte";

  const BASE = import.meta.env.VITE_APP_BASE_URL;
  let recentLogs: string = "Loading logs...";
  let error: string | null = null;

  onMount(() => {
    const ws = new WebSocket(`ws://${BASE}/api/db-logs`);

    ws.addEventListener("open", () => {
      console.log("Connected to ws://" + BASE + "/api/db-logs");
    });

    ws.addEventListener("message", (e) => {
      try {
        const data = JSON.parse(e.data);
        if (data.recent_logs) {
          recentLogs = data.recent_logs;
        } else if (data.error) {
          error = data.error;
        }
      } catch {
        error = "Failed to parse log data";
      }
    });

    ws.addEventListener("error", () => {
      error = "WebSocket error";
    });

    return () => {
      ws.close();
    };
  });
</script>

{#if error}
  <div class="font-bold text-red-600">{error}</div>
{:else}
  <pre
    class="mt-4 max-h-[450px] overflow-scroll rounded-xl font-mono text-sm break-words whitespace-pre-wrap text-gray-800">
    <code class="">{recentLogs}</code>
  </pre>
{/if}
