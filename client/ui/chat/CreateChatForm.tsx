
const CreateChatForm = ({ handleSubmitChatForm, setShowCreateChatForm }) => (
  <form onSubmit={handleSubmitChatForm}>
    <input type="text" placeholder="Chat Name" name="title" required />
    <button type="submit">Create</button>
    <button type="button" onClick={() => setShowCreateChatForm(false)}>Cancel</button>
  </form>
);

export default CreateChatForm;
