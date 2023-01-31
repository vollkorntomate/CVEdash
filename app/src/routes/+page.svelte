<script>
	import { onMount } from 'svelte';
	import CVESummaryView from '../lib/CVESummaryView.svelte';

	let latestPage = 1;
	let latestCVEs = [];

	onMount(fetchLatest);

	async function fetchLatest() {
		await fetch('http://localhost:8077/latest/' + latestPage)
			.then((r) => r.json())
			.then((data) => {
				latestCVEs = latestCVEs.concat(data);
			});
	}

	async function loadMoreLatest() {
		latestPage += 1;
		await fetchLatest();
	}
</script>

<div class="container mx-auto py-4">
	<div class="flex flex-row flex-wrap">
		<div
			class="md:basis-1/3 lg:basis-1/4 order-2 md:order-1 max-h-screen-padded overflow-y-scroll no-scrollbar bg-neutral-100 dark:bg-neutral-800 rounded-lg drop-shadow-2xl"
		>
			{#each latestCVEs as cve}
				<CVESummaryView
					cveID={cve.id}
					cveDescription={cve.description}
					published={cve.published}
					cvssScore={cve.cvssBaseScore}
					cvssSeverity={cve.cvssBaseSeverity}
					cvssVector={cve.cvssVector}
					cvssSource={cve.cvssSource}
				/>
			{/each}
			<div class="text-center">
				<button on:click={loadMoreLatest} class="mb-2 p-2 bg-blue-500 rounded-md">Load more</button>
			</div>
		</div>

		<div class="md:basis-1/3 lg:basis-1/2 order-1 md:order-2">
			<div class="box">DIAGRAM</div>
		</div>

		<div class="md:basis-1/3 lg:basis-1/4 order-3">
			<CVESummaryView
				cveID="twitter"
				cveDescription=""
				published=""
				cvssScore={0.0}
				cvssSeverity=""
				cvssVector=""
				cvssSource=""
			/>
			<CVESummaryView
				cveID="mastodon"
				cveDescription=""
				published=""
				cvssScore={0.0}
				cvssSeverity=""
				cvssVector=""
				cvssSource=""
			/>
			<CVESummaryView
				cveID="and so forth..."
				cveDescription="But of course not as CVESummaryViews..."
				published=""
				cvssScore={0.0}
				cvssSeverity=""
				cvssVector=""
				cvssSource=""
			/>
		</div>
	</div>
</div>
