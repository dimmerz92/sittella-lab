<script>
	import { onMount } from 'svelte';
	import { typewriter } from '$lib';

	let i = -1;
	let text = ['Custom Solutions Crafted from the Ground Up.'];
	let visibility = 'opacity-0';

	onMount(() => {
		const interval = setInterval(() => {
			i += 1;
			i %= text.length;
		}, 500);

		setTimeout(() => {
			visibility = 'opacity-100';
		}, 3000);

		return () => {
			clearInterval(interval);
		};
	});
</script>

<section
	id="home"
	class="flex flex-col my-4 h-52 sm:h-80 4xl:h-[520px] justify-between scroll-m-64"
>
	{#key i}
		<h1
			class=" text-[1.75rem] sm:text-clamp-sm xl:text-clamp-xl text-center sm:text-left mb-2 sm:mb-0 sm:w-[80%] xl:w-[60%]"
			in:typewriter={{ speed: 2 }}
		>
			{text[i] || ''}
		</h1>
	{/key}

	<p class={`${visibility} sm:p-4 transition-opacity ease-in duration-500`}>
		Freelance Web Development and Data Analytics, based in Perth, WA.
	</p>
</section>
