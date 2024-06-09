import React, {useState} from 'react';
import styles from '@/styles/CreateChatForm.module.css'; // 새로운 CSS 파일을 가져옵니다
interface CreateChatFormProps {
  handleSubmitChatForm: (e: React.FormEvent<HTMLFormElement>) => void;
  setShowCreateChatForm: (show: boolean) => void;
}

const CreateChatForm = ({ handleSubmitChatForm, setShowCreateChatForm }: CreateChatFormProps) => {
  const [isPrivateRoom, setIsPrivateRoom] = useState(false);

  const handleRoomTypeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setIsPrivateRoom(e.target.value === 'private');
  };

  return (
    <form onSubmit={handleSubmitChatForm} className={styles.form}>
      <input
        type="text"
        placeholder="Chat Name"
        name="title"
        required
        className={styles.input}
      />
      <div className={styles.roomTypeContainer}>
        <label>
          <input
            type="radio"
            name="roomType"
            value="public"
            defaultChecked
            onChange={handleRoomTypeChange}
          />
          Public Room
        </label>
        <label>
          <input
            type="radio"
            name="roomType"
            value="private"
            onChange={handleRoomTypeChange}
          />
          Private Room
        </label>
      </div>
      {isPrivateRoom && (
        <>
            <input
              type="password"
              placeholder="Password"
              name="password"
              className={styles.input}
            />
        </>
      )}
      <div className={styles.buttonContainer}>
        <button type="submit">Create</button>
        <button type="button" onClick={() => setShowCreateChatForm(false)}>Cancel</button>
      </div>
    </form>
  );
};

export default CreateChatForm;
