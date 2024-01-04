<script>
	let links = [
		{
			href: '#home',
			label: 'Home'
		},
		{
			href: '#about',
			label: 'About'
		},
		{
			href: '#services',
			label: 'Services'
		},
		{
			href: '#contact',
			label: 'Contact'
		}
	];
</script>

<nav class="self-center my-1 sm:ml-auto">
	<ul class="flex justify-center gap-8">
		{#each links as link}
			<li>
				<a href={link.href} class="font-semibold hover:text-sit-orange">{link.label}</a>
			</li>
		{/each}
	</ul>
</nav>
