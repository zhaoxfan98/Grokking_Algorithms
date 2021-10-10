package java;

public class Binarysearch {

	public static void main(String[] args) {
		int arr[]= {2,4,6,9,11,25,37};
		int index=binarysearch(arr,0,arr.length,11);
		if(index==-1) {
			System.out.println("None");

		}else {
	    	System.out.println("Index:"+index);
		}
	}
	public static int binarysearch(int []arr,int left,int right,int value) {
		if(left>right) {
			return -1;
		}
		int mid=(left+right)/2;
		int midvalue=arr[mid];
		if(value>midvalue) {
			return binarysearch(arr,mid+1,right,value);
		}else if(value<midvalue){
			return binarysearch(arr,left,mid-1,value);
		}else {
			return mid;
		}
   }
}
