<script>
	import CVESummaryView from "../lib/CVESummaryView.svelte";
	import {CVESummary} from "../lib/types"
</script>

<div class="container mx-auto my-4">
	<div class="grid gap-4 grid-cols-4">
		<div class="grid auto-rows-auto grid-cols-1 gap-2">
			<CVESummaryView cveSummary="{new CVESummary("CVE-2023-12345", "This CVE is a small test for the box in which a description will live. This CVE is a small test for the box in which a description will live. This CVE is a small test for the box in which a description will live.", "gestern", "LOW", 1.3)}" />
			<CVESummaryView cveSummary="{new CVESummary("asdf", "nice desc", "gestern", "MEDIUM", 4.3)}" />
			<CVESummaryView cveSummary="{new CVESummary("asdf", "nice desc", "gestern", "HIGH", 7.3)}" />
			<CVESummaryView cveSummary="{new CVESummary("asdf", "nice desc", "gestern", "CRITICAL", 9.3)}" />
			<CVESummaryView cveSummary="{new CVESummary("asdf", "nice desc", "gestern", "NO CVSS", 0)}" />
		</div>
		
		<div class="col-span-2">
			<div class="box">DIAGRAM</div>
		</div>
		
		<div class="grid auto-rows-auto grid-cols-1 gap-2">
			<CVESummaryView cveSummary="{new CVESummary("", "twitter", "", "", 0)}" />
			<CVESummaryView cveSummary="{new CVESummary("", "mastodon", "", "", 0)}" />
			<CVESummaryView cveSummary="{new CVESummary("", "und so weiter", "", "", 0)}" />
		</div>
	</div>
</div>