# JS for Blockchain Developers

## Introduction to CSS
- CSS Diner: https://flukeout.github.io/
- Flexbox: https://forum.ivanontech.com/t/assignment-flexbox/41746
- Codesandbox source code of the web layout:
  - https://codesandbox.io/s/strange-bogdan-xzds2 
- Paul Irish CSS Reset: https://www.paulirish.com/2012/box-sizing-border-box-ftw/ 
- Hamburger menu implementation: https://codesandbox.io/s/exciting-goldstine-qy4g5 
  - Source of inspiration for the hamburger menu: https://codepen.io/andreykrokhin/pen/mGEqja 

- Notice that the fastest way to learn is to search codepen.io or CSS websites for solutions you’d like to implement, and re-implement these examples yourself. For instance, you can learn a lot of CSS selector tricks in the hamburger menu example.

## Fundamentals JS
- To understand JavaScript values in depth, it helps if you to read about this topic from another source.
  - Eloquent JavaScript is a great resource, it is definitely worth giving it a read. Don’t forget to solve the exercises at the end of those sections that contain them. You may need to use Google to solve some of the exercises.
  https://eloquentjavascript.net/01_values.html
  - If after searching in Google, some questions are still troubling you with ‘Values, Types, and Operators’ section in the book, please post your questions here:
  https://forum.ivanontech.com/t/reading-assignment-types-and-operators/41719
- Bottom-Up Planning Assigment:
  - Forum link: https://forum.ivanontech.com/t/assignment-bottom-up-planning/41721
  - Complete the Tic-Tac-Toe game by implementing each missing function in the code below. Some variables are not declared, as they are only part of our planning.
  - Feel free to use the Codesandbox link containing the below code: https://codesandbox.io/s/gifted-khayyam-chkqf?file=/game.js. Codesandbox also gives you sufficient error highlighting to signal where undeclared variables are.
  ```js
  function getUserInput(nextPlayerSymbol) {

  }

  function isMoveValid(coordinates, gameBoard) {

  }

  function makeAMove(gameBoard, nextPlayerSymbol) {
      // clone the game board before placing moves in it
      do {
          let coordinates = getUserInput(nextPlayerSymbol);
      } while ( !isMoveValid(coordinates, gameBoard) );
      // return newGameBoard;
  }

  function hasLastMoverWon(lastMove, gameBoard) {
      let winnerCombos = [
          [0, 1, 2], 
          [3, 4, 5], 
          [6, 7, 8],
          [0, 3, 6],
          [1, 4, 7], 
          [2, 5, 8],
          [0, 4, 8], 
          [2, 4, 6]
      ];
      for (let [i1, i2, i3] of winnerCombos) {
          if (gameBoard[i1] === lastMove &&
              gameBoard[i1] === gameBoard[i2] && 
              gameBoard[i1] === gameBoard[i3] 
             ) {
              return true;
          }
      }
      return false;
  }

  function isGameOver(gameBoard, currentPlayerSymbol) {
      // 1. check if there is a winner 
      if (hasLastMoverWon(lastMove, gameBoard) ) {
          // Write a message that last mover has won the game
          alert(`Congratulations, ${currentPlayerSymbol} has won the game`);
          return true;
      }
      // 2. check if the board is full

      // Return: winner/draw OR game is still in progress
  }

  function ticTacToe() {
      // State space 
      let gameBoard = new Array(9).fill(null);
      let players = ['X', 'O'];
      let nextPlayerIndex = 0;

      // Computations 
     do {
          gameBoard = makeAMove(gameBoard, currentPlayerSymbol);
      } while ( !isGameOver(gameBoard, currentPlayerSymbol) );

      // Return value 
      // return undefined;
  } 
  ```

  - Exercise Solution:
  ```js
  function printBoard(gameBoard) {
    let gameString = '';
    for (let i = 0; i <= 6; i += 3) {
        gameString += `${gameBoard[i] ?? i+1}${gameBoard[i+1] ?? i+2}${gameBoard[i+2] ?? i+3}\n`;
    }
    return gameString;
  }

  function getUserInput(nextPlayerSymbol, gameBoard) {
      return +prompt(`Board:\n${printBoard(gameBoard)} Enter your next move (1-9)   :`);
  }

  function isMoveValid(coordinates, gameBoard) {
      return (
          typeof coordinates === 'number' && 
          coordinates >= 1 &&
          coordinates <= 9 && 
          gameBoard[coordinates - 1] === null
      ); 
  }

  function makeAMove(gameBoard, nextPlayerSymbol) {
      const newGameBoard = JSON.parse(JSON.stringify(gameBoard));
      let coordinates;
      do {
          coordinates = getUserInput(nextPlayerSymbol, gameBoard);
      } while ( !isMoveValid(coordinates, gameBoard) );
      newGameBoard[coordinates - 1] = nextPlayerSymbol;
      return newGameBoard;
  }

  function hasLastMoverWon(lastMove, gameBoard) {
      let winnerCombos = [
          [0, 1, 2], 
          [3, 4, 5], 
          [6, 7, 8],
          [0, 3, 6],
          [1, 4, 7], 
          [2, 5, 8],
          [0, 4, 8], 
          [2, 4, 6]
      ];
      for (let [i1, i2, i3] of winnerCombos) {
          if (gameBoard[i1] === lastMove &&
              gameBoard[i1] === gameBoard[i2] && 
              gameBoard[i1] === gameBoard[i3] 
             ) {
              return true;
          }
      }
      return false;
  }

  function isGameOver(gameBoard, currentPlayerSymbol) {
      // 1. check if there is a winner
      if (hasLastMoverWon(currentPlayerSymbol, gameBoard) ) {
          // Write a message that last mover has won the game
          alert(`Congratulations, ${currentPlayerSymbol} has won the game`);
          return true;
      }
      // 2. check if the board is full
      const isBoardFull = gameBoard.every(element => element !== null);
      if (isBoardFull) {
          alert(`${printBoard(gameBoard)}\nDraw. Game Over.`);
          return true;
      }
      return false;
  }

  function ticTacToe() {
      // State space
      let gameBoard = new Array(9).fill(null);
      let players = ['X', 'O'];
      let nextPlayerIndex = 1;

      // Computations
     do {
          nextPlayerIndex = (nextPlayerIndex + 1) % 2;
          currentPlayerSymbol = players[nextPlayerIndex];
          gameBoard = makeAMove(gameBoard, currentPlayerSymbol);
      } while ( !isGameOver(gameBoard, currentPlayerSymbol) );
  } 
  ```

- TicTacToe Code
  - The code https://codesandbox.io/s/optimistic-mclean-csd30?file=/tictactoe.js helps you with the initial state of the application, containing some basic styles for the Tic-Tac-Toe board. You will see how these styles are developed in the next video. Feel free to use this markup or write your own. 
  - The first step of developing an event listener driven game is to remove the console-based game. Therefore, either comment out or delete the ticTacToe() function call in line 105.
  - You can edit this code by registering a Codesandbox account and then making some changes to these files. This copies the original Codesandbox code to your account and saves all your future changes.
  - You can access the final solution with DOM manipulation below: https://codesandbox.io/s/laughing-feather-g8nqo. 
  - While the code solves the Tic Tac Toe problem, a last challenge awaits you: try to win the game with either X or O, and then check out what happens if you click on an empty cell after winning the game. This behavior symbolizes a very frequent phenomenon while developing software: edge cases have to be dealt with. 
  - Think about how you can prevent this behavior to occur in the code, and write your solution in the forum. If you type directly in this Codesandbox with your registered account, then you fork this Sandbox, meaning that the above code will be copied, and your changes will be preserved in your account.

## Asynchronous Programming
- 1inch API Documentation: https://docs.1inch.io/docs/aggregation-protocol/api/swagger/

## DEX
- Check out Ivan’s video on the swap method of the 1inch plugin here: https://www.youtube.com/watch?v=stRh5Scd8TY
- DEX Final Code: https://github.com/Ivan-on-Tech-Academy/js-prog-course-2021
- If you want to learn how to build more complicated dApps and expand your abilities as a blockchain developer then you have to check out Moralis YouTube channel:
  https://www.youtube.com/c/MoralisWeb3
  https://forum.ivanontech.com/t/congratulations-on-completing/41724