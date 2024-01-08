<script>
	import Contact from './contact.svelte';
</script>

<div class="flex flex-col gap-2 pb-2 items-center bg-sit-lblue">
	<Contact />
	<div id="socials" class="flex gap-2">
		<a href="https://www.linkedin.com/company/sittella-lab" class="flex flex-col items-center">
			<img
				class="bg-white rounded-md"
				src="/logos/linkedin.svg"
				alt="Linkedin logo"
				height="25"
				width="25"
			/>
		</a>

		<a href="https://github.com/dimmerz92" class="flex flex-col items-center">
			<img
				class="bg-white rounded-full"
				src="/logos/facebook.svg"
				alt="Facebook logo"
				height="25"
				width="25"
			/>
		</a>

		<a href="https://github.com/dimmerz92" class="flex flex-col items-center">
			<img
				class="bg-white rounded-full"
				src="/logos/github.svg"
				alt="Github logo"
				height="25"
				width="25"
			/>
		</a>
	</div>
	<div class="text-sm">Sittella Lab &#169; 2024. All rights reserved.</div>
</div>
