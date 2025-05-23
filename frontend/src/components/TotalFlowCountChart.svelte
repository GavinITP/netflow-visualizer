<script lang="ts">
  import { onMount } from "svelte";
  import type { ChartConfiguration } from "chart.js";
  import {
    Chart,
    BarController,
    BarElement,
    CategoryScale,
    LinearScale,
    Title,
    Tooltip,
    Legend,
  } from "chart.js";

  Chart.register(
    BarController,
    BarElement,
    CategoryScale,
    LinearScale,
    Title,
    Tooltip,
    Legend,
  );

  export let packetHistory: { time: string; count: number }[];

  const data: ChartConfiguration<"bar">["data"] = {
    labels: packetHistory.map((p) => p.time),
    datasets: [
      {
        label: "Flow Count",
        data: packetHistory.map((p) => p.count),
        backgroundColor: "rgba(75, 192, 192, 0.5)",
        borderColor: "rgba(75, 192, 192, 1)",
        borderWidth: 1,
      },
    ],
  };

  const config: ChartConfiguration<"bar"> = {
    type: "bar",
    data,
    options: {
      responsive: true,
      maintainAspectRatio: false,
      animation: false,
      plugins: {
        title: {
          display: true,
          text: "Recent Flow Count (10 Seconds)",
        },
        legend: {
          display: false,
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
            text: "Flow Count",
          },
        },
      },
    },
  };

  let canvas: HTMLCanvasElement;
  let chart: Chart<"bar">;

  onMount(() => {
    const ctx = canvas.getContext("2d");
    if (ctx) {
      chart = new Chart(ctx, config);
    }
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
