# Comment signals

The comment signals gives more information as compare to MoSCoW comments. Think of them as conventional comments.
Read more about here: [Conventional Comments](https://conventionalcomments.org/).

| Signal           | Meaning / Intent                                            | Example Use Case                                | Example Comment                                                                       |
| ---------------- | ----------------------------------------------------------- | ----------------------------------------------- | ------------------------------------------------------------------------------------- |
| **nitpick**      | Very small suggestion, non-blocking.                        | Formatting, naming, or style preference.        | `nitpick: Could we rename this variable to be more descriptive?`                      |
| **suggestion**   | Non-blocking idea, reviewer proposes an improvement.        | Alternative code snippet, better readability.   | `suggestion: You might consider using a map here instead of a loop for clarity.`      |
| **question**     | Asking for clarification, not necessarily a change request. | Reviewer doesn’t understand intent.             | `question: Why do we need this extra check here?`                                     |
| **praise**       | Positive reinforcement, highlighting good practices.        | Code that’s elegant, clean, or well-tested.     | `praise: Love how concise this helper function is.`                                   |
| **blocker**      | Must be addressed before merging.                           | Security bug, incorrect logic, failing test.    | `blocker: This input isn’t sanitized, which could cause XSS.`                         |
| **major**        | Important issue, needs change but not a merge blocker.      | Performance concern, poor design choice.        | `major: This approach may cause O(n^2) complexity, consider optimizing.`              |
| **minor**        | Small issue, not critical.                                  | Variable naming, small refactor.                | `minor: Variable could be more descriptive.`                                          |
| **typo**         | Grammar or spelling correction.                             | Docs, comments, variable names.                 | `typo: Fix spelling of "occurred".`                                                   |
| **needs change** | Explicit request for change, stronger than "suggestion".    | Wrong implementation, missed requirement.       | `needs change: This function doesn’t handle null input.`                              |
| **rework**       | Larger changes required, could mean restructuring.          | Wrong approach, large refactor.                 | `rework: This service layer mixes concerns, should separate API from business logic.` |
| **clarify**      | Code is unclear, request explanation or docs.               | Complex regex, math formula, nested logic.      | `clarify: Can you add a comment explaining this regex?`                               |
| **blocking**     | Same as **blocker**, but sometimes softer tone.             | Must resolve before merge.                      | `blocking: Unit tests are missing for this new feature.`                              |
| **info**         | Reviewer provides extra context or learning resource.       | Sharing best practices, docs links.             | `info: FYI, you could also achieve this with async/await (see link).`                 |
| **optional**     | Suggestion that’s nice-to-have.                             | Not necessary but helpful.                      | `optional: Extract this into a helper if you think it makes sense.`                   |
| **style**        | Formatting or consistency issue.                            | Indentation, linting, code style guide.         | `style: Use single quotes to match project convention.`                               |
| **security**     | Highlight security-related issues.                          | SQL injection, XSS, unsafe eval.                | `security: Avoid string concatenation in SQL queries.`                                |
| **perf**         | Performance-related feedback.                               | Inefficient algorithm, unnecessary computation. | `perf: This loop recalculates values that could be cached.`                           |
| **UX**           | User experience concern.                                    | Error messages, confusing flow.                 | `UX: Consider showing a toast instead of a silent failure.`                           |
| **test**         | Suggest adding/changing tests.                              | Missing coverage, weak assertions.              | `test: Add a test case for empty input.`                                              |
| **docs**         | Documentation-related.                                      | Missing comments, README updates.               | `docs: Update README with setup instructions.`                                        |
| **future**       | Not needed now but good to keep in mind.                    | Tech debt, scalability, roadmap.                | `future: This will need sharding if dataset grows beyond X.`                          |
