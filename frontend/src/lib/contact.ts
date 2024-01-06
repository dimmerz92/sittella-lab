import { get, post } from '../helpers/http';

type TextResponse = {
	message: string;
};

export async function sendContactMessage(data: FormData) {
	try {
		const payload = Object.fromEntries(data.entries());

		const res: TextResponse = await post(`${import.meta.env.VITE_API_URL}/api/contact`, payload);

		if (!res.message) throw new Error(`${res}`);

		return res.message;
	} catch (err) {
		console.log(err);
	}
}

export async function getContactTemplate() {
	try {
		const res: TextResponse = await get(`${import.meta.env.VITE_API_URL}/api/contact`);

		if (!res.message) throw new Error(`${res}`);

		return res.message;
	} catch (err) {
		console.log(err);
	}
}
