<script lang="ts">
	import { fly } from 'svelte/transition';
	import { inview, type ObserverEventDetails } from 'svelte-inview';

	export let className: string;

	let isInView: boolean;
	const handleChange = ({ detail }: CustomEvent<ObserverEventDetails>) => {
		isInView = detail.inView;
	};
</script>

<div use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }} on:inview_change={handleChange}>
	{#if isInView}
		<div class={className} in:fly={{ y: 200, duration: 1000 }}>
			<slot />
		</div>
	{/if}
</div>
