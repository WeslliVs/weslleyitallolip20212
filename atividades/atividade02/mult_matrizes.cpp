#include <iostream>

using namespace std;

int main(){

  int matriz_a[3][2]{{1, 2}, {3, 4}, {5, 6}};
  int matriz_b[2][3]{{4, 5, 6}, {7, 8, 9}};
  int produto_aeb[3][3] = {{0, 0, 0}, {0, 0, 0}, {0, 0, 0}};

  for (int linha = 0; linha < 3; linha++){
    for (int coluna = 0; coluna < 3; coluna++){
      for (int contador = 0; contador < 2; contador++)
        produto_aeb[linha][coluna] += matriz_a[linha][contador] * matriz_b[contador][coluna];
    }
  }

  for (int linha = 0; linha < 3; linha++)
    for (int coluna = 0; coluna < 3; coluna++)
      cout << produto_aeb[linha][coluna] << endl;

  return 0;
}