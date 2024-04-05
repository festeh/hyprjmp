<script lang="ts">
	import { Autocomplete } from '@skeletonlabs/skeleton';
	import { afterUpdate } from 'svelte';
  import { isNumber } from '$lib';
	export let data;

	let selectionMode = true;

	let sessions = data.props.data.Sessions;
	let input = '';
	let visibleOptions = [...options];

	function onSelection(event) {
		input = event.detail.label;
	}

	afterUpdate(() => {
		console.log(visibleOptions);
	});

	function myFilter() {
		// Create a local copy of options
		let _options = [...options];
		// Filter options
		visibleOptions = _options.filter((option) => {
			// Format the input and option values
			const inputFormatted = String(input).toLowerCase().trim();
			const optionFormatted = option.value.toLowerCase().trim();
			// Check that all letters of the input are included in the option in the same order
			let curOption = optionFormatted;
			for (let i = 0; i < inputFormatted.length; i++) {
				if (!curOption.includes(inputFormatted[i])) {
					return;
				}
				curOption = curOption.slice(curOption.indexOf(inputFormatted[i]));
			}
			return option;
		});
		return visibleOptions;
	}

	let inputCallback = (event) => {
		selectionMode = false;
		document.getElementById('input').focus();
		event.preventDefault();
	};

	let escapeCallback = () => {
		selectionMode = true;
		document.getElementById('input').blur();
	};

	function listener(event) {
		if (event.key === 'i') {
			inputCallback(event);
		} else if (event.key === 'Escape') {
			escapeCallback();
		}
		if (isNumber(event.key)) {
			const key = parseInt(event.key);
			if (key === 0) {
				return;
			}
			if (visibleOptions.length >= key) {
				const opt = visibleOptions[key - 1];
				const reqData = {
					type: 'tmux',
					Session: opt.value.split('|')[0].split(' ')[1],
					Window: opt.value.split('|')[1].split(':')[0]
				};
				console.log(JSON.stringify(reqData));
				fetch('http://localhost:8999/jump', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						Accept: 'application/json'
					},
					body: JSON.stringify(reqData)
				});
			}
		}
	}

	function onKeypress(_node) {
		window.addEventListener('keydown', listener);
		return {
			destroy() {
				window.removeEventListener('keydown', listener);
			}
		};
	}
</script>

<div use:onKeypress>
	<div class="flex w-full">
		<input
			id="input"
			bind:value={input}
			type="search"
			class="input-bordered input m-2 p-1"
			placeholder="Search"
		/>
		<span class="mr-2 w-16 self-center border p-2 text-center font-bold">
			{#if selectionMode}
				S
			{:else}
				I
			{/if}
		</span>
	</div>
	<div class="card m-2 h-full w-full overflow-y-auto p-1" tabindex="-1">
		<Autocomplete bind:input {options} on:selection={onSelection} filter={myFilter} />
	</div>
</div>
