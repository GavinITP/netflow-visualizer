<script lang="ts">
  import { onMount } from "svelte";
  import {
    Chart,
    PieController,
    ArcElement,
    Tooltip,
    Legend,
    Title,
  } from "chart.js";

  let canvas: HTMLCanvasElement;

  onMount(() => {
    Chart.register(PieController, ArcElement, Tooltip, Legend, Title);

    const ctx = canvas.getContext("2d");

    if (ctx) {
      const chart = new Chart(ctx, {
        type: "pie",
        data: {
          labels: ["TCP", "UDP"],
          datasets: [
            {
              data: [29, 71],
              backgroundColor: ["#A593E0", "#D6C1F0"],
              borderColor: "#fff",
              borderWidth: 2,
            },
          ],
        },
        options: {
          responsive: true,
          plugins: {
            title: {
              display: true,
              font: {
                size: 18,
              },
            },
            legend: {
              position: "bottom",
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
