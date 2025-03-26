<script lang="ts">
  import { onMount } from "svelte";
  import {
    Chart,
    LineController,
    LineElement,
    PointElement,
    LinearScale,
    CategoryScale,
    Title,
    Tooltip,
    Legend,
  } from "chart.js";

  let canvas: HTMLCanvasElement;

  onMount(() => {
    Chart.register(
      LineController,
      LineElement,
      PointElement,
      LinearScale,
      CategoryScale,
      Title,
      Tooltip,
      Legend,
    );

    const ctx = canvas.getContext("2d");

    if (ctx) {
      const chart = new Chart(ctx, {
        type: "line",
        data: {
          labels: [
            "11:00:00",
            "11:00:01",
            "11:00:02",
            "11:00:03",
            "11:00:04",
            "11:00:05",
            "11:00:06",
            "11:00:07",
            "11:00:08",
            "11:00:09",
          ],
          datasets: [
            {
              label: "Packets",
              data: [12, 19, 3, 5, 0, 0, 0, 1, 1, 2],
              borderColor: "#9494ff",
              tension: 0.2,
              fill: false,
            },
          ],
        },
        options: {
          responsive: true,
          plugins: {
            title: {
              display: true,
              text: "Anomaly Packets Over Time",
            },
          },
          scales: {
            x: {
              title: {
                display: true,
                text: "Time (HH:MM:SS)",
              },
            },
            y: {
              beginAtZero: true,
              title: {
                display: true,
                text: "Packet Count",
              },
            },
          },
        },
      });

      return () => chart.destroy();
    }
  });
</script>

<div class="flex h-full w-full justify-center">
  <canvas bind:this={canvas} class="h-full w-full"></canvas>
</div>
