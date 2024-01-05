<script lang="ts">
	import { sendGPTMessage, type GPTMessage } from '$lib';
	import { onMount } from 'svelte';
	import { linear } from 'svelte/easing';
	import { slide } from 'svelte/transition';
	import Message from './message.svelte';

	let messages: GPTMessage[] = [];
	let initialised = false;
	let isVisible = false;

	const sendChat = async (e: Event) => {
		e.preventDefault();

		const textArea = document.getElementById('content') as HTMLTextAreaElement;
		if (!initialised) {
			textArea.addEventListener('keypress', (e: KeyboardEvent) => {
				if (e.key === 'Enter' && !e.shiftKey) {
					e.preventDefault();
					(document.getElementById('chat_button') as HTMLButtonElement).click();
				}
			});
		}

		console.log(textArea);

		// extract form data to json
		const content = textArea.value;
		textArea.value = '';

		const message: GPTMessage = {
			role: 'user',
			content: content
		};

		messages = [...messages, message];

		const reply = await sendGPTMessage(messages);

		if (!reply) return;

		messages = [...messages, reply];
	};

	const toggleChat = async () => {
		if (!messages.length)
			do {
				const reply = await sendGPTMessage([]);
				if (!reply) break;
				messages = [...messages, reply];
			} while (false);
		isVisible = !isVisible;
	};
</script>

<div class="flex flex-col gap-2 fixed bottom-0 right-0 w-full sm:w-2/5 xl:w-1/4">
	{#if isVisible}
		<div
			transition:slide={{ duration: 300, easing: linear }}
			class="flex flex-col h-[70vh] z-1000 sm:mr-2 bg-sit-yellow rounded-md px-2"
		>
			<div class="flex justify-center items-center h-10">Virtual Assistant</div>
			<div id="chat_window" class="flex flex-col gap-2 grow p-2 bg-white rounded-md">
				{#each messages as message}
					<Message role={message.role} content={message.content} />
				{/each}
			</div>
			<div id="chat" class="flex gap-2 py-2">
				<textarea
					name="content"
					id="content"
					class="grow rounded-md resize-none box-border"
					rows="3"
					required
				></textarea>
				<button
					id="chat_button"
					on:click={sendChat}
					class="bg-sit-dblue text-white aspect-square rounded-md"
				>
					Send
				</button>
			</div>
		</div>
	{/if}

	<button
		type="submit"
		on:click={toggleChat}
		class="bg-sit-dblue rounded-full p-1.5 ml-auto mr-2 mb-2"
	>
		<img src="/logos/message.svg" alt="message icon" class="h-16" />
	</button>
</div>
