<script lang="ts">
  export let tableData: Array<any> = [];

  // Pagination state
  let currentPage = 1;
  let rowsPerPage = 15;

  $: totalPages = Math.ceil(tableData.length / rowsPerPage);
  $: paginatedData = tableData.slice(
    (currentPage - 1) * rowsPerPage,
    currentPage * rowsPerPage,
  );

  function nextPage() {
    if (currentPage < totalPages) currentPage++;
  }

  function prevPage() {
    if (currentPage > 1) currentPage--;
  }
</script>

<div class="mt-8 overflow-x-auto">
  <table class="min-w-full text-sm">
    <thead class="border-b border-gray-200">
      <tr>
        <th class="px-4 py-2 text-left font-medium text-gray-700">SrcIP</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">SrcPort</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">DstIP</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">DstPort</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">Protocol</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">NextHop</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">dPkts</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">dOctets</th>
      </tr>
    </thead>
    <tbody>
      {#each paginatedData as row}
        <tr class="border-b border-gray-200 hover:bg-gray-50">
          <td class="px-4 py-2 text-gray-800">{row.srcaddr}</td>
          <td class="px-4 py-2 text-gray-800">{row.srcport}</td>
          <td class="px-4 py-2 text-gray-800">{row.dstaddr}</td>
          <td class="px-4 py-2 text-gray-800">{row.dstport}</td>
          <td class="px-4 py-2 text-gray-800">{row.prot}</td>
          <td class="px-4 py-2 text-gray-800">{row.nexthop}</td>
          <td class="px-4 py-2 text-gray-800">{row.dPkts}</td>
          <td class="px-4 py-2 text-gray-800">{row.dOctets}</td>
        </tr>
      {/each}
    </tbody>
  </table>

  <div class="mt-6 flex justify-center">
    <div
      class="inline-flex items-center space-x-2 rounded-lg border border-gray-300 bg-white px-4 py-2 shadow-sm"
    >
      <button
        on:click={prevPage}
        class="rounded-md px-3 py-1 text-sm font-medium text-gray-700 hover:bg-gray-100 disabled:opacity-40"
        disabled={currentPage === 1}
      >
        ‹ Prev
      </button>

      <span class="text-sm text-gray-700">
        Page <strong>{currentPage}</strong> of <strong>{totalPages}</strong>
      </span>

      <button
        on:click={nextPage}
        class="rounded-md px-3 py-1 text-sm font-medium text-gray-700 hover:bg-gray-100 disabled:opacity-40"
        disabled={currentPage === totalPages}
      >
        Next ›
      </button>
    </div>
  </div>
</div>
