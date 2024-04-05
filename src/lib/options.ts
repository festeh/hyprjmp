type TmuxSession = {
  Name: string;
  Windows: TmuxWindow[];
};

type TmuxWindow = {
  Id: number;
  Name: string;
};

export function createTmuxOptions(sessions: TmuxSession[]) {
	const options = sessions.flatMap((session) =>
		session.Windows.map((window) => {
			const name = `(tmux) ${session.Name}| ${window.Id}:${window.Name}`;
			return { value: name, label: name };
		})
	);
  return options;
}
