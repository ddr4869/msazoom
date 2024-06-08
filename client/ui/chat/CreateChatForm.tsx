import React from 'react';
import styles from '@/styles/CreateChatForm.module.css'; // 새로운 CSS 파일을 가져옵니다
interface CreateChatFormProps {
  handleSubmitChatForm: (e: React.FormEvent<HTMLFormElement>) => void;
  setShowCreateChatForm: (show: boolean) => void;
}

const CreateChatForm = ({ handleSubmitChatForm, setShowCreateChatForm }:CreateChatFormProps) => (
  <form onSubmit={handleSubmitChatForm} className={styles.form}>
    <input type="text" placeholder="Chat Name" name="title" required className={styles.input} />
    <div className={styles.buttonContainer}>
      <button type="submit">Create</button>
      <button type="button" onClick={() => setShowCreateChatForm(false)}>Cancel</button>
    </div>
  </form>
);

export default CreateChatForm;
