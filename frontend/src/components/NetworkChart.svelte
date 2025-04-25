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

  export let packetHistory: { time: string; count: number }[] = [];
  let canvas: HTMLCanvasElement;
  let chart: Chart;

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
    if (!ctx) return;

    chart = new Chart(ctx, {
      type: "line",
      data: {
        labels: packetHistory.map((p) => p.time),
        datasets: [
          {
            label: "Packets",
            data: packetHistory.map((p) => p.count),
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
          tooltip: { enabled: true },
          legend: { display: false },
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
              text: "Packet Rate",
            },
          },
        },
      },
    });
  });

  $: if (chart) {
    chart.data.labels = packetHistory.map((p) => p.time);
    chart.data.datasets[0].data = packetHistory.map((p) => p.count);
    chart.update();
  }
</script>

<div class="flex h-full w-full items-center justify-center">
  <canvas bind:this={canvas} class="h-full w-full"></canvas>
</div>
