import React, { useState } from "react";
import { CreateProblem } from "../wailsjs/go/main/App";

function ProblemForm() {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [message, setMessage] = useState("");

    const submitProblem = async () => {
        try {
            await CreateProblem({ title, description });
            setMessage("問題を登録しました！");
            setTitle("");
            setDescription("");
        } catch (err: any) {
            setMessage("エラーが発生しました：" + err.message);
        }
    };

    return (
        <div>
            <h2>問題登録フォーム</h2>
            <input
                type="text"
                placeholder="タイトル"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
            />
            <br />
            <textarea
                placeholder="説明"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
            />
            <br />
            <button onClick={submitProblem}>登録</button>
            <p>{message}</p>
        </div>
    );
}

export default ProblemForm;
