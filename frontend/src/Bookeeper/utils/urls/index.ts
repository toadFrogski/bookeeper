const urls = {
  home: "/",
  signIn: "/sign-in",
  signUp: "/sign-up",
  profile: "/profile",
  profileAddBook: "/profile/book/add",
  profileEditBook: "/profile/book/:bookID",
};

export const parseURL = (path: string): string[] => {
  return path
    .split(/(?=\/)|(?=:)\//g) /* split by path components */
    .filter((e) => e) /* filter empty strings */
};

export default urls;
