import ReactionsComponent from './reactionsComponent';
import '../app/reactionPost/postStyles.scss';

const PostPage: React.FC = () => {
  return (
    <div className="post">
      <img src="https://thumbs.dreamstime.com/z/%C3%A1rbol-de-navidad-19159101.jpg?w=992" alt="Publicación" className="post-image" />
      <ReactionsComponent />
    </div>
  );
};

export default PostPage;