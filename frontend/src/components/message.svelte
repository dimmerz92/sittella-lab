<script lang="ts">
	export let role: string;
	export let content: string;
</script>

<div
	class={`w-fit p-1 rounded-md ${
		role === 'user' ? 'bg-sit-lblue ml-auto text-white' : 'bg-slate-200'
	}`}
>
	{content}
</div>
