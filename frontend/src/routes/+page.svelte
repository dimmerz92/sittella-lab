<script>
	import Footer from '../components/footer.svelte';
	import Hero from '../components/hero.svelte';
	import LoadOnScroll from '../components/loadOnScroll.svelte';
	import PreferredTech from '../components/preferredTech.svelte';

	const duration = 1500;
	const margin = '-100px';
	const y = 200;
</script>

<Hero />

<div class="flex flex-col gap-12 sm:gap-24 my-24">
	<LoadOnScroll {duration} {y} {margin}>
		<div class="flex flex-col gap-2 bg-slate-200 sm:w-[60%] p-4 sm:ml-auto rounded-md">
			<h2 class="text-xl">Digital Solutions, <span class="text-sit-orange">Simplified.</span></h2>
			<p class="text-justify">
				I can build your website or help you make sense of your data. Allow me to make your digital
				journey easier with simple solutions for a range of requirements without the complex tech
				jargon.
			</p>
		</div>
	</LoadOnScroll>

	<section id="about" class="scroll-m-64 rounded-md">
		<LoadOnScroll {duration} {y} {margin}>
			<div class="flex flex-col sm:flex-row items-center bg-slate-200 p-4 gap-8 rounded-md">
				<div class="flex flex-col gap-8 sm:w-[60%]">
					<div class="flex flex-col gap-2 text-justify">
						<h2 class="text-xl">About <span class="text-sit-orange">Sittella</span></h2>
						<p class="font-semibold">Sittella</p>
						<p><em>Noun /sɪˈtɛlə/</em></p>
						<p>
							A Sittella is a tiny, nimble, and agile Australian bird, expertly building it's nest
							among the branches
						</p>
						<p>
							Sittella Lab embodies the same precision and craftsmanship in the digital realm. Not
							content with being "just a developer", Sittella Lab aims to build seemless, robust,
							user-friendly, and accessible web solutions.
						</p>
					</div>
					<div class="flex flex-col gap-2 text-justify">
						<h2 class="text-xl">About <span class="text-sit-orange">Andrew</span></h2>
						<p>
							Hi, I'm Andrew, a web developer and data enthusiast with postgraduate studies in
							computer science and data science. As the mind behind Sittella Lab, I specialise in
							turning ideas into digital solutions. With a foundation in data analytics, I bring
							precision and efficiency to web development. Let's build something powerful together.
						</p>
					</div>
				</div>
				<div class="relative sm:w-[40%]">
					<div
						class="absolute top-1/2 transform -translate-y-1/2 right-[10%] sm:right-[15%] 4xl:right-1/4 w-min p-2 bg-slate-200 bg-opacity-75 rounded-md"
					>
						<p class="text-2xl sm:text-4xl text-sit-dblue font-semibold">Hi, I'm Andrew.</p>
					</div>
					<img src="/images/andrew.jpg" alt="Andrew: Owner of Sittella Lab" class="rounded-md" />
				</div>
			</div>
		</LoadOnScroll>
	</section>

	<section id="services" class="scroll-m-64">
		<LoadOnScroll {duration} {y} {margin}>
			<div
				class="flex flex-col sm:flex-row-reverse sm:items-center bg-slate-200 p-4 gap-8 rounded-md"
			>
				<div class="flex flex-col gap-8 bg-slate-200 sm:w-[60%] p-4 rounded-md">
					<h2 class="text-xl">Services:</h2>
					<div>
						<h3 class="text-sit-orange">Full-Stack Web Development</h3>
						<ul class="flex flex-col gap-2 sm:text-justify">
							<li class="list-disc ml-4">
								<strong>Responsiveness:</strong> Seemless experiences over a variety of platforms and
								devices
							</li>
							<li class="list-disc ml-4">
								<strong>End-to-End:</strong> From intuitive and interactive user interfaces, powerful
								and dynamic server logic, and database infrastructure
							</li>
							<li class="list-disc ml-4">
								<strong>Security & Efficiency:</strong> Prioritising website and user security and application
								performance
							</li>
						</ul>
					</div>
					<div>
						<h3 class="text-sit-orange">Data Analytics</h3>
						<ul class="flex flex-col gap-2 sm:text-justify">
							<li class="list-disc ml-4">
								<strong>Data Insights & Interpretation:</strong> Gain valuable insights with data analysis
								and interpretation of your data
							</li>
							<li class="list-disc ml-4">
								<strong>Visualise:</strong> Transform your data into clear and easy to understand visuals
								for any audience
							</li>
							<li class="list-disc ml-4">
								<strong>Strategic Data Driven Decisions:</strong> Empower your business with actionable
								data insights to drive strategy and decision making
							</li>
						</ul>
					</div>
				</div>
				<div class="sm:w-[40%]">
					<PreferredTech />
				</div>
			</div>
		</LoadOnScroll>
	</section>
</div>
