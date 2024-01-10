<div class="spinner"></div>

<style>
	.spinner {
		border: 5px solid theme(colors.sit-dblue);
		border-top: 5px solid theme(colors.sit-orange);
		border-radius: 50%;
		width: 25px;
		height: 25px;
		animation: spin 2s linear infinite;
	}

	@keyframes spin {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
