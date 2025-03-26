<script lang="ts">
  import { onMount } from "svelte";
  // @ts-ignore
  import hljs from "highlight.js/lib/core";

  const logLang = (hljs: any) => ({
    name: "log",
    contains: [
      { className: "section", begin: /\[\w+\]/ },
      { className: "string", begin: /"/, end: /"/ },
      { className: "number", begin: /\b\d+(\.\d+)?(ms|s|d)?\b/ },
      {
        className: "keyword",
        begin:
          /\b(SELECT|FROM|WHERE|AND|OR|INSERT|DELETE|UPDATE|DROP|POST|HTTP)\b/,
        case_insensitive: true,
      },
      { className: "comment", begin: /\/[^\s"]+/ },
    ],
  });

  let code = `[httpd] 2024-03-12T12:30:45Z ::1 - - [12/Mar/2024:12:30:45 +0000] "POST /write?db=mydb HTTP/1.1" 204 0 "-" "InfluxDBClient" 12abc-78-defg5ms

[query] 2024-03-12T12:31:10Z SELECT * FROM temperature WHERE time > now() - 1h

[store] 2024-03-12T12:32:00Z Compacted 2 TSM files to /var/lib/influxdb/data/mydb/autogen/1234.tsm

[retention] 2024-03-12T12:33:15Z Retention policy enforced - dropped 1000 points older than 30d

[error] 2024-03-12T12:34:00Z Error writing point: field type conflict`;

  onMount(() => {
    hljs.registerLanguage("log", logLang);
    hljs.highlightAll();
  });
</script>

<div>
  <pre
    class="whitespace-pre-wrap break-words rounded-xl font-mono text-sm text-gray-800">
    <code class="language-log">{code}</code>
  </pre>
</div>
