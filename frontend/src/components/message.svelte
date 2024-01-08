<script lang="ts">
	export let role: string;
	export let content: string;
</script>

<div
	class={`w-fit max-w-[70%] p-4 rounded-md text-xs ${
		role === 'user' ? 'bg-sit-lblue ml-auto text-white' : 'bg-slate-200'
	}`}
>
	{content}
</div>
