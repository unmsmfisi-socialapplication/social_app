This database includes tables for user accounts, profiles, friendships, posts, comments, reactions and interests. It is designed to store data related to user interactions and content shared within the social network, making it a comprehensive database for managing various aspects of a social networking platform.

# Tables descriptionw

<table>
  <thead>
    <tr>
      <th>Tabla</th>
      <th>Descripción</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>SOC_APP_USERS</td>
      <td>This table will store the accounts registered in the platform.</td>
    </tr>
    <tr>
      <td>SOC_APP_USER_PROFILE</td>
      <td>This table will store the profile information of the registered accounts.</td>
    </tr>
    <tr>
      <td>SOC_APP_FRIENDS</td>
      <td>This table will store the friends of the registered accounts.</td>
    </tr>
    <tr>
      <td>SOC_APP_M_USERS_INTERESTS</td>
      <td>This table will be a master that will contain all the interests that will be held within the social network.</td>
    </tr>
    <tr>
      <td>SOC_APP_POSTS</td>
      <td>All posts made by users will be stored in this table.</td>
    </tr>
    <tr>
      <td>SOC_APP_POSTS_COMMENTS</td>
      <td>All comments related to a post will be stored in this table.</td>
    </tr>
    <tr>
      <td>SOC_APP_POSTS_REACTIONS</td>
      <td>In this table all the reactions related to a post will be stored.</td>
    </tr>
    <tr>
      <td>SOC_APP_POSTS_COMMENTS_REACTIONS</td>
      <td>All reactions related to a comment will be stored in this table.</td>
    </tr>
    <tr>
      <td>SOC_APP_USERS_INTERESTS_POSTS</td>
      <td>This table will store the interests associated with a publication.</td>
    </tr>
    <tr>
      <td>SOC_APP_M_USERS_REACTIONS</td>
      <td>This table will be a master that will contain all the types of reactions that will be available on the social network.</td>
    </tr>
  </tbody>
</table>

# ERD Database



![Image](https://github.com/unmsmfisi-socialapplication/social_app/assets/67833630/6007808f-b091-468c-a12b-5f15ba86bcd2)