# Post Interview Thoughts

- There was a bug in the code: We were modifying ErrCounter incorrectly because the constructors of LLMProviders were implemented incorrectly.
- Constructors of LLMProviders should have returned address to the object instead of the object itself. Because this was implemented incorrectly, we kept on seeing errorCounter = 0 even tho requests were failing.