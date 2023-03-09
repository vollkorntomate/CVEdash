<script>
	let licenses = [
		{
			name: 'Chart.js',
			license: 'MIT',
			url: 'https://github.com/chartjs/Chart.js/blob/master/LICENSE.md'
		},
		{
			name: 'chartjs-plugin-datalabels',
			license: 'MIT',
			url: 'https://github.com/chartjs/chartjs-plugin-datalabels/blob/master/LICENSE.md'
		},
		{
			name: 'ForkAwesome',
			license: 'MIT',
			url: 'https://github.com/ForkAwesome/Fork-Awesome/blob/master/LICENSES'
		},
		{
			name: 'Moment.js',
			license: 'MIT',
			url: 'https://github.com/moment/moment/blob/develop/LICENSE'
		},
		{
			name: 'SvelteKit',
			license: 'MIT',
			url: 'https://github.com/sveltejs/kit/blob/master/LICENSE'
		},
		{
			name: 'Tailwind CSS',
			license: 'MIT',
			url: 'https://github.com/tailwindlabs/tailwindcss/blob/master/LICENSE'
		},
		{
			name: 'tw-elements',
			license: 'License',
			url: 'https://github.com/mdbootstrap/Tailwind-Elements/blob/master/license'
		}
	];
</script>

<div class="dashboard-card p-4 py-2 w-full">
	<div class="overflow-y-scroll no-scrollbar h-full">
		<div class="my-2">
			<p class="text-lg">
				This site uses the NVD API, but is not endorsed or certified by the NVD.
			</p>
		</div>

		<hr class="my-2" />

		<div class="my-2">
			<p class="text-lg">
				This site uses open source software, whose individual licenses can be found below.
			</p>
		</div>

		<ul class="my-2 list-disc list-inside">
			{#each licenses as license}
				<li>
					<span class="font-bold text-lg">{license.name}</span>
					<span>(<a href={license.url}>{license.license}</a>)</span>
				</li>
			{/each}
		</ul>
	</div>
</div>

<style>
	a {
		text-decoration: underline;
	}
</style>
