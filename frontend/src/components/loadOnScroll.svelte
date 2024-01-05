<script lang="ts">
	import { fly } from 'svelte/transition';
	import { inview, type ObserverEventDetails } from 'svelte-inview';

	export let duration: number;
	export let margin: string;
	export let y: number;

	let isInView: boolean;
	const handleChange = ({ detail }: CustomEvent<ObserverEventDetails>) => {
		isInView = detail.inView;
	};
</script>

<div use:inview={{ unobserveOnEnter: true, rootMargin: margin }} on:inview_change={handleChange}>
	{#if isInView}
		<div in:fly={{ y: y, duration: duration }}>
			<slot />
		</div>
	{/if}
</div>
