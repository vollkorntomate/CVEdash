<script lang="ts">
	import moment from 'moment';
	import type { CveSummary } from '$lib/types';

	export let cve: CveSummary;

	$: color = '';
	$: switch (cve.cvssBaseSeverity) {
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
			<a href="https://nvd.nist.gov/vuln/detail/{cve.id}" target="_blank" rel="noreferrer noopener">
				<span class="text-xl font-bold xl:whitespace-nowrap">{cve.id}</span>
			</a>
		</div>
		<div class="text-right">
			<div>
				{#if cve.cvssBaseScore > 0.0 && cve.cvssBaseSeverity !== ''}
					<span
						class="p-1 text-sm rounded-md {color}"
						title="Vector: {cve.cvssVector}&#013;Source: {cve.cvssSource}"
					>
						{cve.cvssBaseScore.toFixed(1)}
						{cve.cvssBaseSeverity}
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
				<span class="text-sm" title={new Date(cve.published).toUTCString()}>
					{moment(cve.published).fromNow()}
				</span>
			</div>
		</div>
	</div>
	<div>
		<p style="overflow-wrap: break-word;">{cve.description}</p>
	</div>
</div>
