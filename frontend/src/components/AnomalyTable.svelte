<script lang="ts">
  let { tableData } = $props();

  // let tableData = [
  //   {
  //     timestamp: "2025-03-26 12:00:00",
  //     srcIp: "192.168.0.1",
  //     dstIp: "10.0.0.2",
  //     protocol: "TCP",
  //     nextHop: "192.168.0.254",
  //     bytes: 1024,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:05:00",
  //     srcIp: "192.168.0.3",
  //     dstIp: "10.0.0.4",
  //     protocol: "UDP",
  //     nextHop: "192.168.0.254",
  //     bytes: 512,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:10:00",
  //     srcIp: "192.168.0.5",
  //     dstIp: "10.0.0.6",
  //     protocol: "TCP",
  //     nextHop: "192.168.0.1",
  //     bytes: 2048,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:15:00",
  //     srcIp: "192.168.0.6",
  //     dstIp: "10.0.0.8",
  //     protocol: "UDP",
  //     nextHop: "192.168.0.1",
  //     bytes: 768,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:20:00",
  //     srcIp: "192.168.0.7",
  //     dstIp: "10.0.0.10",
  //     protocol: "ICMP",
  //     nextHop: "192.168.0.1",
  //     bytes: 256,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:25:00",
  //     srcIp: "192.168.0.8",
  //     dstIp: "10.0.0.12",
  //     protocol: "TCP",
  //     nextHop: "192.168.0.254",
  //     bytes: 1536,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:30:00",
  //     srcIp: "192.168.0.9",
  //     dstIp: "10.0.0.14",
  //     protocol: "UDP",
  //     nextHop: "192.168.0.254",
  //     bytes: 640,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:35:00",
  //     srcIp: "192.168.0.10",
  //     dstIp: "10.0.0.16",
  //     protocol: "TCP",
  //     nextHop: "192.168.0.1",
  //     bytes: 3072,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:40:00",
  //     srcIp: "192.168.0.11",
  //     dstIp: "10.0.0.18",
  //     protocol: "ICMP",
  //     nextHop: "192.168.0.1",
  //     bytes: 128,
  //   },
  //   {
  //     timestamp: "2025-03-26 12:45:00",
  //     srcIp: "192.168.0.12",
  //     dstIp: "10.0.0.20",
  //     protocol: "UDP",
  //     nextHop: "192.168.0.254",
  //     bytes: 896,
  //   },
  // ];
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
