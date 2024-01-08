<div class="loading"></div>

<style>
	.loading {
		position: relative;
		width: 10px;
		height: 10px;
		border-radius: 5px;
		background-color: theme(colors.sit-orange);
		color: theme(colors.sit-orange);
		animation: loading 1s infinite linear alternate;
		animation-delay: 0.5s;
	}

	.loading::before,
	.loading::after {
		content: '';
		display: inline-block;
		position: absolute;
		top: 0;
	}

	.loading::before {
		left: -15px;
		width: 10px;
		height: 10px;
		border-radius: 5px;
		background-color: theme(colors.sit-orange);
		color: theme(colors.sit-orange);
		animation: loading 1s infinite alternate;
		animation-delay: 0s;
	}

	.loading::after {
		left: 15px;
		width: 10px;
		height: 10px;
		border-radius: 5px;
		background-color: theme(colors.sit-orange);
		color: theme(colors.sit-orange);
		animation: loading 1s infinite alternate;
		animation-delay: 1s;
	}

	@keyframes loading {
		0% {
			background-color: theme(colors.sit-orange);
		}
		50%,
		100% {
			background-color: rgba(152, 128, 255, 0.2);
		}
	}
</style>
