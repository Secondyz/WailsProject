interface Window {
    backend: {
        App: {
            CreateProblem: (input: { title: string; description: string }) => Promise<void>;
            Greet: (name: string) => Promise<string>;
        };
    };
}
