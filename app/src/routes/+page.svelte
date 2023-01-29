<script>
	import { CVESummary } from "$lib/types";
	import { onMount } from "svelte";
	import CVESummaryView from "../lib/CVESummaryView.svelte";

	let latestCVEs = [];

	onMount(async () => {
		await fetch('http://localhost:8077/latest/1')
		.then(r => r.json())
		.then(data => {
			latestCVEs = data
		});
	});
</script>

<div class="container mx-auto my-4">
	<div class="grid gap-4 grid-cols-4">
		<div class="grid auto-rows-auto grid-cols-1 gap-2">
			{#each latestCVEs as cve}
				<CVESummaryView cveID="{cve.id}" cveDescription="{cve.description}" published="{cve.published}" cvssScore="{cve.cvssBaseScore}" cvssSeverity="{cve.cvssBaseSeverity}" />
			{/each}
		</div>
		
		<div class="col-span-2">
			<div class="box">DIAGRAM</div>
		</div>
		
		<div class="grid auto-rows-auto grid-cols-1 gap-2">
			<CVESummaryView cveID="twitter" cveDescription="" published="" cvssScore="{0.0}" cvssSeverity="" />
			<CVESummaryView cveID="mastodon" cveDescription="" published="" cvssScore="{0.0}" cvssSeverity="" />
			<CVESummaryView cveID="and so forth..." cveDescription="But of course not as CVESummaryViews..." published="" cvssScore="{0.0}" cvssSeverity="" />
		</div>
	</div>
</div>