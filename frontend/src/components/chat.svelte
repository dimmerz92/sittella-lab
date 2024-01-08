<script lang="ts">
	import { sendGPTMessage, type GPTMessage } from '$lib';
	import { linear } from 'svelte/easing';
	import { slide } from 'svelte/transition';
	import Message from './message.svelte';
	import Loading from './loading.svelte';
	import { onMount } from 'svelte';

	let messages: GPTMessage[] = [];
	let textArea: HTMLTextAreaElement;
	let initialised = false;
	let isVisible = false;
	let isLoading = false;
	let lastEnter: Date | undefined;
	const minTime = 2000; // milliseconds

	onMount(() => {
		textArea = document.getElementById('content') as HTMLTextAreaElement;
	});

	const toggleChat = async () => {
		isVisible = !isVisible;
		if (!initialised) await initialiseChat();
	};

	const initialiseChat = async () => {
		isLoading = true;

		const reply = await sendGPTMessage([]);
		if (!reply) return;

		messages = [...messages, reply];

		initialised = true;
		isLoading = false;
	};

	const alterEnter = (e: KeyboardEvent) => {
		if (e.key === 'Enter' && !e.shiftKey) {
			e.preventDefault();
			if (textArea && textArea.value.length < 3) return;

			//prevent spamming
			if (!lastEnter || (lastEnter && new Date().getTime() - lastEnter.getTime() >= minTime)) {
				(document.getElementById('chat_button') as HTMLButtonElement).click();
				lastEnter = new Date();
			}
		}
	};

	const sendMessages = async (e: Event) => {
		e.preventDefault();

		// get current user message and append to array
		const message: GPTMessage = {
			role: 'user',
			content: textArea.value
		};
		messages = [...messages, message];
		textArea.value = '';

		// send the updated array
		isLoading = true;
		const reply = await sendGPTMessage(messages);
		if (!reply) return;

		// update the array with response
		messages = [...messages, reply];
		isLoading = false;
	};
</script>

<div class="fixed bottom-0 right-0 max-h-full max-w-full flex flex-col justify-end items-end z-20">
	<div
		transition:slide={{ duration: 300, easing: linear }}
		class={`${
			isVisible ? 'h-screen sm:h-[768px] p-2' : 'h-0'
		} flex flex-col bg-sit-yellow w-screen sm:w-96 gap-2 sm:mr-4 rounded-md shadow-lg overflow-hidden`}
	>
		<div class="relative flex justify-center items-center h-10">Virtual Assistant</div>
		<div
			id="chat_window"
			class="relative flex flex-col grow bg-white rounded-md gap-2 p-2 overflow-scroll"
		>
			{#each messages as message}
				<Message role={message.role} content={message.content} />
			{/each}
			{#if isLoading}
				<div class="self-center mt-auto">
					<Loading />
				</div>
			{/if}
		</div>
		<div id="chat" class="flex gap-2">
			<textarea
				name="content"
				id="content"
				on:keypress={alterEnter}
				class="grow resize-none rounded-md"
				rows="3"
				required
			></textarea>
			<button
				id="chat_button"
				on:click={sendMessages}
				class="bg-sit-dblue text-white aspect-square rounded-md"
			>
				Send
			</button>
		</div>
	</div>

	<button type="submit" on:click={toggleChat} class="bg-sit-dblue p-1 m-2 rounded-full">
		<img src="/logos/message.svg" alt="message icon" class="h-16" />
	</button>
</div>
