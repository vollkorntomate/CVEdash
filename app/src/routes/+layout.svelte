<script>
	import '../app.postcss';
</script>

<div class="md:h-screen overflow-scroll md:overflow-hidden no-scrollbar">
	<div class="flex flex-col h-full">
		<nav class="relative w-full flex flex-wrap items-center justify-between py-2 bg-nav dark:bg-nav-dark text-text1-dark">
			<div class="container-fluid w-full flex flex-wrap items-center justify-between px-6 text-xl">
				<div>
					<a href="/">
						<span class="font-bold">CVEdash</span>
					</a>
				</div>
				<div>
					<a href="https://github.com/vollkorntomate/CVEdash" target="_blank" rel="nofollow noreferrer" class="ml-4">
						<span class="text-2xl fa fa-github" title="See source on GitHub"></span>
					</a>
					<a href="legal" class="ml-4">
						<span class="fa fa-balance-scale" title="Legal notes"></span>
					</a>
				</div>
			</div>
		</nav>

		<div class="container flex mx-auto px-2 py-4 min-h-0">
			<slot />
		</div>
	</div>
</div>