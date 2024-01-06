<script lang="ts">
	import { getContactTemplate, sendContactMessage } from '$lib';
	import { onMount } from 'svelte';

	let formContainer: HTMLElement;
	let formText: HTMLTextAreaElement;

	onMount(() => {
		formContainer = document.getElementById('form_container')!;
		formText = document.getElementById('message') as HTMLTextAreaElement;
	});

	const getTemplate = async (e: Event) => {
		e.preventDefault();

		const reply = await getContactTemplate();
		if (!reply) return;

		formText.value += reply;
	};

	const submitForm = async (e: Event) => {
		e.preventDefault();

		const form = (e.target as HTMLElement).closest('form');
		if (!form) throw new Error('target is not a form');

		const formData = new FormData(form);

		const reply = await sendContactMessage(formData);
		if (!reply) return;

		formContainer.style.height = '60px';
		formContainer.classList.remove(...['h-full']);
		formContainer.textContent = reply;
	};
</script>

<div
	id="form_container"
	class="formContainer flex flex-col gap-4 items-center bg-sit-yellow h-[480px] w-5/6 sm:w-1/2 4xl:w-1/4 p-4 rounded-b-md transition-all ease-in-out duration-1000"
>
	<div class="flex flex-col gap-2">
		<h2 class="text-2xl font-medium">Get in touch.</h2>
		<p>If you have a design ready to go, or you just want to have a chat, drop us a message!</p>
		<p>
			Not sure what to ask or say? Add this <button class="text-sit-orange" on:click={getTemplate}
				>template</button
			> to your enquiry.
		</p>
	</div>
	<form id="contact" class="flex flex-col gap-2 w-full" on:submit={submitForm}>
		<div class="flex gap-1 h-8">
			<label for="first_name" class="sr-only">First Name</label>
			<input
				type="text"
				id="first_name"
				class="grow min-w-0 px-2 rounded-md"
				name="first_name"
				placeholder="* First Name"
				required
			/>

			<label for="last_name" class="sr-only">Last Name</label>
			<input
				type="text"
				id="last_name"
				class="grow min-w-0 px-2 rounded-md"
				name="last_name"
				placeholder="* Last Name"
				required
			/>
		</div>
		<label for="email" class="sr-only">Email</label>
		<input
			type="email"
			id="email"
			class="h-8 px-2 rounded-md"
			name="email"
			placeholder="* Email"
			required
		/>

		<label for="message" class="sr-only">Message</label>
		<textarea
			id="message"
			class="resize-none px-2 rounded-md"
			rows="5"
			name="message"
			placeholder="* Message"
			required
		/>
		<button type="submit" class="text-sit-orange hover:text-sit-dblue text-lg font-medium">
			Send Message
		</button>
	</form>
</div>
