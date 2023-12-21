<script>
	import { AppShell, Autocomplete } from '@skeletonlabs/skeleton';
	export let data;

	let selectionMode = true;

	let sessions = data.props.data.Sessions;
	let input = '';
	const options = sessions.flatMap((session) =>
		session.Windows.map((window) => {
			const name = `(tmux) ${session.Name}| ${window.Id}:${window.Name}`;
			return { value: name, label: name };
		})
	);

	function onSelection(event) {
		input = event.detail.label;
	}

	function myFilter() {
		// Create a local copy of options
		let _options = [...options];
		// Filter options
		return _options.filter((option) => {
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
			class="input input-bordered m-2 p-1"
			placeholder="Search"
		/>
		<span class="w-16 self-center p-2 text-center border font-bold mr-2">
			{#if selectionMode}
				S
			{:else}
				I
			{/if}
		</span>
	</div>
	<div class="card w-full h-full p-1 m-2 overflow-y-auto" tabindex="-1">
		<Autocomplete bind:input {options} on:selection={onSelection} filter={myFilter} />
	</div>
</div>
