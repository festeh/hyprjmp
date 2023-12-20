<script>
export let data;

let sessions = data.props.data.Sessions;
console.log(JSON.stringify(sessions));
</script>
<ul>
{#each sessions as session}
   {#each session.Windows as window}
      <li>
      <span>
      (tmux) {session.Name}| {window.Id}:{window.Name}
      </span>
      </li>
   {/each}
{/each}
</ul>
