export function typewriter(el: HTMLElement, { speed = 1 }: { speed: number }) {
	if (!el.textContent) return {};

	const text = el.textContent;
	const duration = text.length / (speed * 0.01);

	return {
		duration,
		tick: (t: number) => {
			const i = Math.trunc(text!.length * t);
			el.textContent = text.slice(0, i);
		}
	};
}
