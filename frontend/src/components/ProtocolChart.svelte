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

  let { protocolData } = $props<{ protocolData: Record<string, number> }>();

  let labels = $derived(Object.keys(protocolData));
  let values = $derived(
    Object.values(protocolData).map((v) => Math.round(Number(v) * 100)),
  );

  const baseColors = ["#A593E0", "#D6C1F0"];

  let colors = $derived(
    labels.map((_, i, arr) =>
      baseColors[i]
        ? baseColors[i]
        : `hsl(${(i * 360) / arr.length}, 70%, 60%)`,
    ),
  );

  let canvas: HTMLCanvasElement;
  let chart: Chart;

  onMount(() => {
    Chart.register(PieController, ArcElement, Tooltip, Legend, Title);
    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    chart = new Chart(ctx, {
      type: "pie",
      data: {
        labels,
        datasets: [
          {
            data: values,
            backgroundColor: colors,
            borderColor: "#fff",
            borderWidth: 2,
          },
        ],
      },
      options: {
        responsive: true,
        plugins: {
          legend: { position: "bottom" },
        },
      },
    });

    return () => chart.destroy();
  });

  $effect(() => {
    if (!chart) return;
    chart.data.labels = labels;
    chart.data.datasets[0].data = values;
    chart.data.datasets[0].backgroundColor = colors;
    chart.update();
  });
</script>

<div class="flex h-full w-full justify-center">
  <canvas bind:this={canvas} class="h-full w-full"></canvas>
</div>
