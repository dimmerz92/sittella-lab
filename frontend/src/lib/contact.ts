import { get, post } from '../helpers/http';

type TextResponse = {
	message: string;
};

export async function sendContactMessage(data: FormData) {
	try {
		const payload = Object.fromEntries(data.entries());
		const res: TextResponse = await post('/api/contact', payload);
		return res.message;
	} catch (err) {
		console.log(err);
	}
}

export async function getContactTemplate() {
	try {
		const res: TextResponse = await get('/api/contact');
		return res.message;
	} catch (err) {
		console.log(err);
	}
}
