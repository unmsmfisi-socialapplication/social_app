
import React from "react";
import PostPage from "@/components/postPage";
import Layout from "../layout";
// import './postStyles.scss'

const PostHome: React.FC = () => {
  return (
    <div className="post-container">
      <h1>Posts</h1>
      <PostPage />
      <PostPage />
      <PostPage />
      <PostPage />
    </div>
  );
};

export default PostHome;