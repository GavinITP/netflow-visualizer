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
    Object.values(protocolData).map((v) => Number(v) * 100),
  );

  const baseColors = ["#A593E0", "#FFC1F0", "#F0B3C7", "#FFDA91"];

  let colors = $derived(
    labels.map((_, i) => baseColors[i % baseColors.length]),
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

<div class="flex h-full w-full justify-center py-5">
  <canvas bind:this={canvas} class="h-full w-full"></canvas>
</div>
