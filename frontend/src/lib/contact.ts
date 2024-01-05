export async function sendContactMessage(data: FormData) {
	try {
		const res = await fetch(`${import.meta.env.VITE_API_URL}/api/chat`, {
			method: 'POST',
			body: data
		});

		if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

		return await res.text();
	} catch (err) {
		console.log(err);
		return '';
	}
}
