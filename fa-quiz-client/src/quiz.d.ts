export type QuizEntry =
    | {
          /**
           * @description entry type for quiz item
           */
          item: "question";

          /**
           * @description whether the question is the current active question
           */
          isActive: boolean;

          /**
           * @pocketbase `questions.id`
           */
          id: string;

          /**
           * @description this quiz belongs to chapter ...
           */
          chapterIndex: number;

          /**
           * @pocketbase `questions.question`
           */
          question: string;

          /**
           * @pocketbase `questions.correctAnswer` + questions.otherAnswers - stored as a comma-separated string in the database
           * @description shuffled, remember where correct answer is after shuffling
           */
          options: string[];
          correctAnswer: number;

          /**
           * @pocketbase `questions.footer`
           */
          footer: string | null;

          /**
           * @description if the quiz has been attempted, which answer if yes?
           */
          selectedAnswers: number[];

          /**
           * @description statistics for the question, only available after quiz completion
           */
          stats: {
              solved: boolean;
              tries: number;
              triesWrong: number;
          };
      }
    | {
          /**
           * @description entry type for quiz item
           */
          item: "definition";

          /**
           * @description whether the definition is the current active quiz entry
           */
          isActive: boolean;

          /**
           * @pocketbase `lecture.id`
           */
          id: string;

          /**
           * @pocketbase `lecture.title`
           */
          subtitle: string;

          /**
           * @pocketbase `lecture.content`
           */
          content: string;
      }
    | {
          /**
           * @description entry type for quiz item
           */
          item: "chapter-finish";

          /**
           * @description whether the definition is the current active quiz entry
           */
          isActive: boolean;

          /**
           * @pocketbase `lecture.id`
           */
          id: string;

          /**
           * @description this quiz belongs to chapter ...
           */
          chapterIndex: number;

          /**
           * @pocketbase `chapters.title`
           */
          title: string;
      };
