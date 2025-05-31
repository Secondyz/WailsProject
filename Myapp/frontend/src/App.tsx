import React, { useState } from "react";
import { CreateProblem } from "../wailsjs/go/problems/ProblemsAPI";
import "./App.css";

function ProblemForm() {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [message, setMessage] = useState("");
    const [isLoading, setIsLoading] = useState(false);

    const submitProblem = async () => {
        if (!title.trim() || !description.trim()) {
            setMessage("タイトルと説明を入力してください");
            return;
        }

        setIsLoading(true);
        try {
            await CreateProblem({ title, description });
            setMessage("問題を登録しました！");
            setTitle("");
            setDescription("");
        } catch (err: any) {
            setMessage("エラーが発生しました：" + err.message);
        } finally {
            setIsLoading(false);
        }
    };

    const handleKeyDown = (e: React.KeyboardEvent) => {
        if (e.key === 'Enter' && (e.ctrlKey || e.metaKey)) {
            submitProblem();
        }
    };

    return (
        <div className="app-container">
            {/* タイトルバー */}
            <div className="title-bar">
                <div className="title-bar-icon"></div>
                <h1 className="title-bar-text">問題管理</h1>
            </div>

            {/* メインコンテンツ */}
            <div className="main-content">
                {/* サイドバー */}
                <div className="sidebar">
                    <div className="sidebar-header">
                        <h2 className="sidebar-title">メニュー</h2>
                    </div>
                    <nav className="nav">
                        <button className="nav-item nav-item-active">
                            新規登録
                        </button>
                        <button className="nav-item nav-item-inactive">
                            問題一覧
                        </button>
                        <button className="nav-item nav-item-inactive">
                            検索
                        </button>
                        <button className="nav-item nav-item-inactive">
                            設定
                        </button>
                    </nav>
                </div>

                {/* メインパネル */}
                <div className="content-area">
                    {/* コンテンツヘッダー */}
                    <div className="content-header">
                        <h2 className="content-title">新しい問題を登録</h2>
                        <p className="content-subtitle">問題の詳細情報を入力してください</p>
                    </div>

                    {/* フォームエリア */}
                    <div className="form-area">
                        <div className="form-container">
                            {/* タイトル */}
                            <div className="form-group">
                                <label className="form-label">
                                    タイトル <span className="required">*</span>
                                </label>
                                <input
                                    type="text"
                                    className="form-input"
                                    placeholder="問題のタイトルを入力"
                                    value={title}
                                    onChange={(e) => setTitle(e.target.value)}
                                    onKeyDown={handleKeyDown}
                                />
                            </div>

                            {/* 説明 */}
                            <div className="form-group">
                                <label className="form-label">
                                    説明 <span className="required">*</span>
                                </label>
                                <textarea
                                    className="form-textarea"
                                    placeholder="問題の詳細な説明を入力してください"
                                    value={description}
                                    onChange={(e) => setDescription(e.target.value)}
                                    onKeyDown={handleKeyDown}
                                />
                            </div>

                            {/* メッセージ */}
                            {message && (
                                <div className={`message ${message.includes('エラー') ? 'message-error' : 'message-success'}`}>
                                    {message}
                                </div>
                            )}
                        </div>
                    </div>

                    {/* ボタンエリア */}
                    <div className="button-area">
                        <div className="shortcut-hint">
                            Ctrl+Enter で登録
                        </div>
                        <div className="button-group">
                            <button
                                className="btn btn-secondary"
                                onClick={() => {
                                    setTitle("");
                                    setDescription("");
                                    setMessage("");
                                }}
                            >
                                クリア
                            </button>
                            <button
                                className={`btn ${isLoading || !title.trim() || !description.trim() ? 'btn-disabled' : 'btn-primary'}`}
                                onClick={submitProblem}
                                disabled={isLoading || !title.trim() || !description.trim()}
                            >
                                {isLoading ? '登録中...' : '登録'}
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            {/* ステータスバー */}
            <div className="status-bar">
                <div>問題管理システム v1.0</div>
                <div>{new Date().toLocaleString('ja-JP')}</div>
            </div>
        </div>
    );
}

export default ProblemForm;