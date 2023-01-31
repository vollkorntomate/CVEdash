<script lang="ts">
	import moment from 'moment';

	export let cveID: string;
	export let cveDescription: string;
	export let published: string;
	export let cvssScore: number;
	export let cvssSeverity: string;
	export let cvssVector: string;
	export let cvssSource: string;

	$: color = '';
	$: switch (cvssSeverity) {
		case '':
			color = 'text-neutral-800 bg-neutral-300 dark:bg-neutral-400';
			break;
		case 'LOW':
			color = 'text-neutral-800 bg-yellow-300';
			break;
		case 'MEDIUM':
			color = 'text-neutral-800 bg-orange-400';
			break;
		case 'HIGH':
			color = 'text-neutral-50 dark:text-neutral-50 bg-red-500 dark:bg-red-600';
			break;
		case 'CRITICAL':
			color = 'text-neutral-100 bg-red-800 dark:bg-red-900';
			break;
	}
</script>

<div class="mb-2 bg-bg3 dark:bg-bg3-dark p-2 rounded-lg">
	<div class="grid grid-cols-2 mb-2">
		<div>
			<a href="https://nvd.nist.gov/vuln/detail/{cveID}" target="_blank" rel="noreferrer noopener">
				<p class="text-xl font-bold xl:whitespace-nowrap">{cveID}</p>
			</a>
		</div>
		<div class="text-right">
			<div>
				{#if cvssScore > 0.0 && cvssSeverity !== ''}
					<span
						class="p-1 text-sm rounded-md {color}"
						title="Vector: {cvssVector}&#013;Source: {cvssSource}"
					>
						{cvssScore.toFixed(1)}
						{cvssSeverity}
					</span>
				{:else}
					<span
						class="p-1 text-sm rounded-md {color}"
						title="CVE has not been scored with CVSS yet"
					>
						NO CVSS
					</span>
				{/if}
			</div>
			<div>
				<span class="text-sm" title={new Date(published).toUTCString()}>
					{moment(published).fromNow()}
				</span>
			</div>
		</div>
	</div>
	<div>
		<p style="overflow-wrap: break-word;">{cveDescription}</p>
	</div>
</div>
