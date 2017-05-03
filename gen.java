package meme;

import java.io.FileNotFoundException;
import java.io.PrintWriter;
import java.util.Random;

public class gen {

	public static void main(String[] args) throws FileNotFoundException{
		Random rand = new Random();
		PrintWriter pw1 = new PrintWriter("set1.txt");
		PrintWriter pw2 = new PrintWriter("set2.txt");
		PrintWriter pw3 = new PrintWriter("test.txt");
		
		pw1.print(rand.nextInt(50) + "," + rand.nextInt(50));	
		
		for(int i = 0; i < rand.nextInt(49); i++){
			pw1.print("," + rand.nextInt(50) + "," + rand.nextInt(50));		
		}		
		
		pw2.print(rand.nextInt(50) + "," + rand.nextInt(50));
		
		for(int i = 0; i < rand.nextInt(49); i++){
			pw2.print(rand.nextInt(50) + "," + rand.nextInt(50) + ",");		
		}
		
		pw3.print(rand.nextInt(50) + "," + rand.nextInt(50));	
		
		for(int i = 0; i < rand.nextInt(49); i++){
			pw3.print(rand.nextInt(50) + "," + rand.nextInt(50) + ",");		
		}		
		
		pw1.close();
		pw2.close();
		pw3.close();		
	}	
}
