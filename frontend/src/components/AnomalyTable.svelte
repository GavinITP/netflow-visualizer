<script lang="ts">
  let { tableData } = $props();
</script>

<div class="mt-8 overflow-x-auto">
  <table class="min-w-full text-sm">
    <thead class="border-b border-gray-200">
      <tr>
        <th class="px-4 py-2 text-left font-medium text-gray-700">Timestamp</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">SrcIP</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">SrcPort</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">DstIP</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">DstPort</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">Protocol</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">NextHop</th>
        <th class="px-4 py-2 text-left font-medium text-gray-700">dPkts</th>
      </tr>
    </thead>
    <tbody>
      {#each tableData as row}
        <tr class="border-b border-gray-200 hover:bg-gray-50">
          <td class="px-4 py-2 text-gray-800">{row.timestamp}</td>
          <td class="px-4 py-2 text-gray-800">{row.srcaddr}</td>
          <td class="px-4 py-2 text-gray-800">{row.srcport}</td>
          <td class="px-4 py-2 text-gray-800">{row.dstaddr}</td>
          <td class="px-4 py-2 text-gray-800">{row.dstport}</td>
          <td class="px-4 py-2 text-gray-800">{row.prot}</td>
          <td class="px-4 py-2 text-gray-800">{row.nexthop}</td>
          <td class="px-4 py-2 text-gray-800">{row.dPkts}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
